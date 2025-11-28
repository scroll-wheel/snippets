export default class Puzzle {
  array: number[];
  rows: number;
  columns: number;

  constructor(rows: number, columns: number) {
    this.rows = rows;
    this.columns = columns;

    this.array = new Array(rows * columns);
    for (let i = 0; i < rows * columns; i++) {
      this.array[i] = i + 1;
    }
  }

  shuffle() {
    do {
      // Fisher–Yates shuffle
      for (let i = this.array.length - 1; i > 0; i--) {
        const j = Math.floor(Math.random() * (i + 1));
        [this.array[i], this.array[j]] = [this.array[j], this.array[i]];
      }
    } while (!this.solvable());
  }

  private solvable(): boolean {
    const array = this.array.filter((n) => n !== this.rows * this.columns);

    let inversions = 0;
    for (let i = 0; i < array.length; i++) {
      for (let j = i + 1; j < array.length; j++) {
        if (array[i] > array[j]) {
          inversions++;
        }
      }
    }

    if (this.columns % 2 === 1) {
      return inversions % 2 === 0;
    }

    const distance =
      this.rows -
      1 -
      Math.floor(this.array.indexOf(this.columns * this.rows) / this.columns);
    console.log(distance);
    return (inversions + distance) % 2 === 0;
  }
}
