import "./style.css";

class Carousel extends HTMLElement {
  content: HTMLOListElement;
  previous: HTMLButtonElement;
  next: HTMLButtonElement;

  constructor() {
    super();
    this.className = "carousel";

    this.content = document.createElement("ol");
    this.content.className = "carousel__content";
    for (let i = 0; i < 5; i++) {
      const item = document.createElement("li");
      item.className = "carousel__item";
      const div = document.createElement("div");
      div.className = "card";
      div.textContent = (i + 1).toString();
      item.appendChild(div);
      this.content.appendChild(item);
    }

    this.previous = (() => {
      const elem = document.createElement("button");
      elem.disabled = true;
      elem.className = "carousel__previous carousel__previous--disabled";

      elem.addEventListener("click", () => {
        if (this.index !== 0) {
          this.index--;
          this.next.disabled = false;
          this.next.classList.remove("carousel__next--disabled");

          if (this.index === 0) {
            elem.disabled = true;
            elem.classList.add("carousel__previous--disabled");
          }
        }
      });

      elem.textContent = "Previous";
      return elem;
    })();

    this.next = (() => {
      const elem = document.createElement("button");
      elem.className = "carousel__next";
      if (this.content.children.length === 1) {
        elem.disabled = true;
        elem.classList.add("carousel__next--disabled");
      }

      elem.addEventListener("click", () => {
        if (this.index !== this.content.children.length - 1) {
          this.index++;
          this.previous.disabled = false;
          this.previous.classList.remove("carousel__previous--disabled");

          if (this.index === this.content.children.length - 1) {
            elem.disabled = true;
            elem.classList.add("carousel__next--disabled");
          }
        }
      });

      elem.textContent = "Next";
      return elem;
    })();

    this.appendChild(this.content);
    this.appendChild(this.previous);
    this.appendChild(this.next);
  }

  get index(): number {
    return +this.getAttribute("index")!;
  }

  set index(index: number) {
    this.setAttribute("index", "" + index);
  }

  static get observedAttributes() {
    return ["index"];
  }

  attributeChangedCallback(name: string) {
    if (name === "index") {
      const scrollLeftMax = this.content.scrollWidth - this.content.clientWidth;
      const width = scrollLeftMax / (this.content.children.length - 1);
      this.content.scrollLeft = this.index * width;
    }
  }
}

customElements.define("my-carousel", Carousel);

document.querySelector<HTMLDivElement>("#app")!.appendChild(new Carousel());
