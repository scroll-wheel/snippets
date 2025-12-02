import matplotlib.pyplot as plt
import numpy as np
import pandas as pd


def hue_vs(field: str):
    plt.title("Saturation vs. Lightness")
    fig = plt.figure()
    ax = fig.add_subplot(projection="polar")
    ax.set_ylim(0, 1)
    c = ax.scatter(
        df["hue"] * np.pi / 180,
        df[field],
        s=(72 / fig.dpi) ** 2,
        c=df["hue"],
        cmap="hsv",
    )
    plt.savefig(f"hue-vs-{field}.png")


if __name__ == "__main__":
    # How can df be read from inside make_fig?
    df = pd.read_csv("data.csv")
    hue_vs("saturation")
    hue_vs("lightness")

    plt.title("Saturation vs. Lightness")
    fig = plt.figure()
    ax = fig.add_subplot()
    ax.set_xlabel("Saturation")
    ax.set_ylabel("Lightness")
    ax.scatter(df["saturation"], df["lightness"], alpha=0.01)
    plt.savefig("saturation-vs-lightness.png")
