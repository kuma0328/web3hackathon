export const BelongingCommunity = () => {
  return (
    <div className="flex gap-10 overflow-x-auto">
      {[0, 1, 2, 3].map(() => {
        return (
          <div className="my-3 h-20 w-20 flex-none rounded-full bg-slate-700"></div>
        );
      })}
    </div>
  );
};
