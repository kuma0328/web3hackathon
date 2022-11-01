import { TypoGraphy } from '../BaseParts/TypoGraphy';

export const RecipeInputs = () => {
  return (
    <div className="lg:p-5">
      <div className="m-3 grid grid-cols-1 lg:m-10">
        <div className="my-5">
          <div className="flex items-center">
            <TypoGraphy className="my-3 text-lg">材料</TypoGraphy>
            <button className="ml-4 rounded-md border border-black py-2 px-5 shadow-lg">
              材料を追加する＋
            </button>
          </div>
          <div className="grid grid-cols-3 place-items-center justify-items-center gap-2 overflow-y-auto"></div>
        </div>

        <div className="my-5">
          <div className="flex items-center">
            <TypoGraphy className="my-3 text-lg">作り方</TypoGraphy>
            <button className="ml-4 rounded-md border border-black py-2 px-5 shadow-lg">
              作り方を追加する＋
            </button>
          </div>
        </div>
      </div>
    </div>
  );
};
