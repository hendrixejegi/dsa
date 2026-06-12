export const duration = (fn) => {
  console.time('Runtime');
  fn();
  console.timeEnd('Runtime');
};
