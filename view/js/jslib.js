function CreateComponent(ele, html) {
    const Templater = document.createElement('template');
    Templater.innerHTML = html;

    class Custom extends HTMLElement {
        constructor() {

            // Always call super first in constructor
            super();
        }

        connectedCallback() {
            const shadowRoot = this.attachShadow({
                mode: 'closed'
            });
            shadowRoot.appendChild(Templater.content);
        }
    }

    customElements.define(ele, Custom);
};