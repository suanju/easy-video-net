export const vRemoveFocus = {
    mounted: (el: any) => {
      el.addEventListener("focus", () => {
        el.blur();
      });
    },
};
  
