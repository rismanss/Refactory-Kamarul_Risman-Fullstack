const list_letters_first = ['c', 'd', 'e', 'g', 'h'];
const list_letters_second = ['X', 'Z'];

const missing = (arr) => {
  for (let i in arr) {
    const letter = arr[i];
    const next = arr[Number(i) + 1];
    if (next.charCodeAt(0) - letter.charCodeAt(0) === 2) {
      return String.fromCharCode(arr[i].charCodeAt(0) + 1);
    }
  }
};

console.log(missing(list_letters_first));
console.log(missing(list_letters_second));
