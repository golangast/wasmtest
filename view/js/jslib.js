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

//https://kinsta.com/blog/web-components/
function CreateFormComponent(ele, html) {

(() => {

  // stop form submission and output field names/values to console
  const form = document.getElementById(ele);

  form.addEventListener('submit', e => {
    
    e.preventDefault();
    console.clear();
    console.log('Submitted data:');

    const data = new FormData(form);
    for (let nv of data.entries()) {
      console.log(`  ${ nv[0] }: ${ nv[1] }`);
    }

  });

})();


// web component
class Custom extends HTMLElement {

  static formAssociated = true;
  
  formAssociatedCallback(form) {
    console.log('form associated:', form.id);
  }
  
  constructor() {
    super();
    this.internals = this.attachInternals();
    this.setValue('');
  }
  
  // connect component
  connectedCallback() {
    
    const shadow = this.attachShadow({ mode: 'closed' });
    shadow.innerHTML = html;
    
    // monitor input values
    shadow.querySelector('input').addEventListener('change', e => {
      this.setValue(e.target.value);
    });
    
  }
  
  // set form value
  setValue(v) {
    this.value = v;
    this.internals.setFormValue(v);
  }
  
}

// register component
customElements.define( ele, Custom );

}