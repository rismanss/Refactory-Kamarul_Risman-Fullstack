const arrayList = [3, 12, 4, 5, 8, 9]

const sortMethod = (arr) => {
  let temp = 0;
  for (let i = 0; i < arr.length; i++) {
    for (let j = i; j < arr.length; j++) {
      if (arr[j] < arr[i]) {
        temp = arr[j];
        arr[j] = arr[i];
        arr[i] = temp;
      }
    }
  }
  return arr;
}

console.log(sortMethod(arrayList));