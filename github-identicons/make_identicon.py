from PIL import Image, ImageDraw

import sys


def make_identicon(
    bits: tuple[
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
        bool,
    ],
    color: tuple[int, int, int],
):
    img = Image.new("RGB", (420, 420), (240, 240, 240))
    draw = ImageDraw.Draw(img)

    for i, fill in enumerate(bits):
        if fill:
            cy = 70 * (i % 5) + 35
            cx = 70 * (i // 5) + 35
            draw.rectangle((cx, cy, cx + 70, cy + 70), fill=color)
            cx = 70 * (4 - i // 5) + 35
            draw.rectangle((cx, cy, cx + 70, cy + 70), fill=color)

    img.save(sys.stdout, "PNG")


if __name__ == "__main__":
    make_identicon(
        (
            # Column 1 & 5
            False,
            True,
            False,
            True,
            False,
            # Column 2 & 4
            True,
            True,
            False,
            True,
            True,
            # Column 3
            False,
            False,
            True,
            False,
            False,
        ),
        (150, 191, 233),
    )
