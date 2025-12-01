package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	fmt.Println("username,r1c1,r2c1,r3c1,r4c1,r5c1,r1c2,r2c2,r3c2,r4c2,r5c2,r1c3,r2c3,r3c3,r4c3,r5c3,red,green,blue,hue,saturation,lightness")

	client := &http.Client{}
	for i := 1; i <= 33000; i++ {
		username := strconv.Itoa(i)

		img, err := func() (image.Image, error) {
			resp, err := client.Get(fmt.Sprintf("https://github.com/identicons/%s.png", username))
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()
			if resp.StatusCode == 404 {
				return nil, fmt.Errorf("Username '%s' does not exist", username)
			}

			img, err := png.Decode(resp.Body)
			if err != nil {
				return nil, fmt.Errorf("Unable to decode response as PNG for username '%s'", username)
			}
			return img, nil
		}()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}

		var red uint8 = 240
		var green uint8 = 240
		var blue uint8 = 240

		fmt.Printf("%s,", username)
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 5; j++ {
				var truthy int
				r, g, b, _ := img.At(70*i, 70*j).RGBA()
				if !(r == 61680 && g == 61680 && b == 61680) {
					red = uint8(r / 256)
					green = uint8(g / 256)
					blue = uint8(b / 256)
					truthy = 1
				}
				fmt.Printf("%d,", truthy)
			}
		}
		hue, saturation, lightness := hsl(red, green, blue)
		fmt.Printf("%d,%d,%d,%f,%f,%f\n", red, green, blue, hue, saturation, lightness)

		time.Sleep(1 * time.Second)
	}
}

// https://www.rapidtables.com/convert/color/rgb-to-hsl.html
func hsl(r, g, b uint8) (float64, float64, float64) {
	rPrime := float64(r) / 255
	gPrime := float64(g) / 255
	bPrime := float64(b) / 255
	cMax := max(rPrime, gPrime, bPrime)
	cMin := min(rPrime, gPrime, bPrime)
	delta := cMax - cMin

	var h float64
	if delta == 0 {
		h = 0
	} else if cMax == rPrime {
		h = 60 * math.Mod((gPrime-bPrime)/delta, 6)
	} else if cMax == gPrime {
		h = 60 * ((bPrime-rPrime)/delta + 2)
	} else if cMax == bPrime {
		h = 60 * ((rPrime-gPrime)/delta + 4)
	}
	if h < 0 {
		h += 360
	}

	var s float64
	if delta != 0 {
		s = delta / (1 - math.Abs(cMax+cMin-1))
	}

	l := (cMax + cMin) / 2

	return h, s, l
}
