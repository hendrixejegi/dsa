import { createArr } from '../../utils/array.js';
import { duration } from '../../utils/duration.js';

function insertionSort(arr) {
  if (arr.length < 2) {
    return;
  }

  for (let i = 1; i < arr.length; i++) {
    const key = arr[i];

    let j = i - 1;
    while (j >= 0 && arr[j] > key) {
      arr[j + 1] = arr[j];
      j--;
    }

    arr[j + 1] = key;
  }
}

(function main() {
  const arr = createArr(10);
  console.log(arr);
  duration(() => insertionSort(arr));
  console.log(arr);
})();
