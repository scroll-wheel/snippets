import './style.css';

import Puzzle from './puzzle';

const tileWidth = 100;
const tileHeight = 100;

class GameTile extends HTMLButtonElement {
  constructor(
    rows: number,
    columns: number,
    n: number,
    row: number,
    column: number
  ) {
    super();
    this.className = 'board__tile';
    this.textContent = n.toString();
    this.style.width = `${tileWidth}px`;
    this.style.height = `${tileHeight}px`;

    this.setAttribute('row', row.toString());
    this.setAttribute('column', column.toString());
    this.addEventListener('click', () => {
      const row = +this.getAttribute('row')!;
      const column = +this.getAttribute('column')!;

      if (
        row > 0 &&
        this.parentElement?.querySelector(
          `.board__tile[row="${row - 1}"][column="${column}"]`
        ) === null
      ) {
        this.setAttribute('row', row - 1 + '');
      } else if (
        row < rows - 1 &&
        this.parentElement?.querySelector(
          `.board__tile[row="${row + 1}"][column="${column}"]`
        ) === null
      ) {
        this.setAttribute('row', row + 1 + '');
      } else if (
        column > 0 &&
        this.parentElement?.querySelector(
          `.board__tile[row="${row}"][column="${column - 1}"]`
        ) === null
      ) {
        this.setAttribute('column', column - 1 + '');
      } else if (
        column < columns - 1 &&
        this.parentElement?.querySelector(
          `.board__tile[row="${row}"][column="${column + 1}"]`
        ) === null
      ) {
        this.setAttribute('column', column + 1 + '');
      }
    });
  }

  static get observedAttributes() {
    return ['row', 'column'];
  }

  attributeChangedCallback(name: string, _: string, newValue: string) {
    if (name === 'row') {
      this.style.top = `${+newValue * tileHeight}px`;
    } else if (name === 'column') {
      this.style.left = `${+newValue * tileWidth}px`;
    }
  }
}

customElements.define('game-tile', GameTile, { extends: 'button' });

class GameBoard extends HTMLElement {
  constructor(puzzle: Puzzle) {
    super();
    this.className = 'board';
    this.style.width = `${puzzle.columns * tileWidth}px`;
    this.style.height = `${puzzle.rows * tileHeight}px`;

    for (let i = 0; i < puzzle.array.length; i++) {
      const tile = puzzle.array[i];
      if (tile !== puzzle.rows * puzzle.columns) {
        this.appendChild(
          new GameTile(
            puzzle.rows,
            puzzle.columns,
            tile,
            Math.floor(i / puzzle.columns),
            i % puzzle.columns
          )
        );
      }
    }

    puzzle.shuffle();
    for (let i = 0; i < puzzle.array.length; i++) {
      if (puzzle.array[i] !== puzzle.rows * puzzle.columns) {
        const index = puzzle.array[i] - 1;
        this.children[index].setAttribute(
          'row',
          Math.floor(i / puzzle.columns).toString()
        );
        this.children[index].setAttribute(
          'column',
          (i % puzzle.columns).toString()
        );
      }
    }
  }
}

customElements.define('game-board', GameBoard);

class An_App extends HTMLElement {
  constructor() {
    super();
    this.appendChild(new GameBoard(new Puzzle(4, 4)));
  }
}

customElements.define('an-app', An_App);

document.querySelector<HTMLDivElement>('#app')!.appendChild(new An_App());
