type TProps = {
  className?: string;
};
export const BelongingCommunity = ({ className }: TProps) => {
  return (
    <div className="flex gap-10 overflow-x-auto">
      {[0, 1, 2, 3].map((t) => {
        return (
          <div
            className={`${className} my-3 h-20 w-20 flex-none bg-slate-700`}
            key={t}
          ></div>
        );
      })}
    </div>
  );
};
