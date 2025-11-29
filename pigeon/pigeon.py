from PIL import Image
from tkinter import Tk
from tkinter.filedialog import askopenfilename

import argparse
import math

def main():
    parser = argparse.ArgumentParser(description="")
    parser.add_argument("-i", "--image", metavar="imagefile", help="")
    parser.add_argument("-p", "--palette", metavar="palettefile", help="", default="base_colors.txt")
    parser.add_argument("-r", "--resize", metavar=("width", "height"), nargs=2, help="")
    parser.add_argument("-s", "--staircase", help="", action="store_true")
    parser.add_argument("-d", "--dithering", help="", action="store_true")
    args = parser.parse_args()

    Tk().withdraw()
    image = args.image if args.image else askopenfilename()
    modify_image(image, args.palette, args.resize, args.staircase, args.dithering)

def modify_image(image, palette, size, staircase, dithering):
    size = (int(size[0]), int(size[1]))
    im = Image.open(image)
    im = im.resize(size)
    pix = im.load()
    width, height = im.size
    table = create_table(palette, staircase)
    resources = dict()

    for y in range(height):
        for x in range(width):
            new_color = get_closest_color(pix[x,y], table)
            quant_error = \
                pix[x,y][0] - new_color["R"], \
                pix[x,y][1] - new_color["G"], \
                pix[x,y][2] - new_color["B"]
            pix[x,y] = new_color["R"], new_color["G"], new_color["B"]

            if dithering:
                dither(pix, x, y, width, height, quant_error)

            if new_color["name"] not in resources:
                resources[new_color["name"]] = 1
            else:
                resources[new_color["name"]] += 1
    
    im.save("preview.png")
    print_resources(resources)

def create_table(palette, staircase):
    table = []
    multipliers = (180 / 255, 220 / 255, 255 / 255)

    for tup in open(palette, "r"):
        for i in range(3):
            data = tup.split(",")
            color = { \
                "id": int(data[0]) * 4 + i, "name": data[1], \
                "R": math.floor(int(data[2]) * multipliers[i]), \
                "G": math.floor(int(data[3]) * multipliers[i]), \
                "B": math.floor(int(data[4]) * multipliers[i])  \
            }
            if staircase or i == 1:
                table.append(color)

    return table

def get_closest_color(pixel, table):
    min_distance = 195075
    closest_color = None

    for color in table:
            distance = \
                (pixel[0] - color["R"]) ** 2 + \
                (pixel[1] - color["G"]) ** 2 + \
                (pixel[2] - color["B"]) ** 2
            if distance < min_distance:
                min_distance = distance
                closest_color = color
                
    return closest_color

def dither(pix, x, y, width, height, quant_error):
    if x + 1 < width:
        pix[x+1,y  ] = pix[x+1,y  ][0] + math.floor(quant_error[0] * 7/16), \
                       pix[x+1,y  ][1] + math.floor(quant_error[1] * 7/16), \
                       pix[x+1,y  ][2] + math.floor(quant_error[2] * 7/16)
    if x - 1 >= 0 and y + 1 < height:
        pix[x-1,y+1] = pix[x-1,y+1][0] + math.floor(quant_error[0] * 5/16), \
                       pix[x-1,y+1][1] + math.floor(quant_error[1] * 5/16), \
                       pix[x-1,y+1][2] + math.floor(quant_error[2] * 5/16)
    if y + 1 < height:
        pix[x  ,y+1] = pix[x  ,y+1][0] + math.floor(quant_error[0] * 3/16), \
                       pix[x  ,y+1][1] + math.floor(quant_error[1] * 3/16), \
                       pix[x  ,y+1][2] + math.floor(quant_error[2] * 3/16)
    if x + 1 < width and y + 1 < height:
        pix[x+1,y+1] = pix[x+1,y+1][0] + math.floor(quant_error[0] * 1/16), \
                       pix[x+1,y+1][1] + math.floor(quant_error[1] * 1/16), \
                       pix[x+1,y+1][2] + math.floor(quant_error[2] * 1/16)

def print_resources(resources):
    for k, v in sorted(resources.items(), key=(lambda x: x[1]), reverse=True):
        print(str(v) + " " + k)

if __name__ == "__main__":
    main()
