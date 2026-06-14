export const createArr = (length) => {
  const arr = [];
  for (let i = 0; i < length; i++) {
    arr[i] = Math.round(Math.random() * 1000);
  }
  return arr;
};
