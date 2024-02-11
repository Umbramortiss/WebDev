const dessertsDropdown = document.getElementById("dessert-hover");
const nav = document.querySelector("nav");
const dessertsAnimation = document.querySelector(".start-desserts");
const dropdownContent = document.querySelector(".dropdown-content");
const addPlaceOrderButtons = document.querySelectorAll('.btn-primary');


// Simplified dropdown animation logic
dessertsDropdown.addEventListener("mouseenter", () => {
  nav.classList.add("dropdown");
  dessertsAnimation.classList.add("active");
});

dessertsDropdown.addEventListener("mouseleave", () => {
  nav.classList.remove("dropdown");
  dessertsAnimation.classList.remove("active");
});

const cart = []; // Initialize an empty array to store cart items

addPlaceOrderButtons.forEach(button => {
  button.addEventListener('click', placeOrder);
});

function placeOrder(event) {
    const productCard = event.target.closest('.card');
    const productName = productCard.querySelector('.card-title').textContent;
    const productPrice = productCard.querySelector('.product-price').textContent;
  
    // Create a new cart item element
    const cartItem = document.createElement('div');
    cartItem.classList.add('cart-icon');
    cartItem.textContent = `${productName} - ${productPrice}`;
  
    // Append the cart item to the cart container
    const cartContainer = document.getElementById('cart-icon');
    cartContainer.appendChild(cartItem);
  
    // Show a notification that the item has been added to the cart
    showNotification(`Added to Cart: ${productName} - ${productPrice}`);
  }

function showNotification(message) {
  // Create a notification element
  const notification = document.createElement('div');
  notification.className = 'notification';
  notification.textContent = message;

  // Append the notification to the body
  document.body.appendChild(notification);

  // Remove the notification after a few seconds (adjust the timeout value as needed)
  setTimeout(() => {
    notification.remove();
  }, 3000); // Notification will disappear after 3 seconds (3000 milliseconds)
}

let cartCount = 0;
function addToCart() {
  cartCount++;
  document.getElementById('cart-count').innerText = cartCount;
  // Display a notification or update UI
  alert('Item added to cart. Total items: ' + cartCount);
  // Additional logic to actually add the item to the cart goes here
}

let timer;
window.onscroll = function() {
  clearTimeout(timer);
  timer = setTimeout(() => {
    const navbar = document.getElementById("navbar");
    if (document.body.scrollTop > 50 || document.documentElement.scrollTop > 50) {
        navbar.classList.add("scrolled");
    } else {
        navbar.classList.remove("scrolled");
    }
  }, 100);
};