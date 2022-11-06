import { TypoGraphy } from '../BaseParts/TypoGraphy';

export const Recipe = () => {
  const contents = ['アジの切り身/サーモン', '酢', '米'];
  const processes = [
    '酢を米と混ぜ合わせる',
    '米をたく',
    '米を整形する',
    '刺身を米の大きさに合わせて切り分ける',
    '米の上に刺身を置く',
  ];
  return (
    <div className="lg:p-5">
      <div className="m-3 grid grid-cols-1 lg:m-10">
        <div>
          <TypoGraphy className="my-3 text-lg">材料</TypoGraphy>
        </div>
        <div className="my-5">
          <div className="grid grid-cols-3 place-items-center justify-items-center gap-2 overflow-y-auto">
            {contents.map((content) => {
              return (
                <div className="w-full rounded-md border border-black p-4 shadow-md">
                  {content}
                </div>
              );
            })}
          </div>
        </div>

        <div className="my-5">
          <TypoGraphy className="my-3 text-lg">作りかた</TypoGraphy>
          <div className="overflow-y-auto">
            {processes.map((process, idx) => {
              return (
                <div className="flex p-2">
                  <div className="mr-10 flex w-10 items-center justify-center rounded-md border border-black p-3 text-center">
                    {idx}
                  </div>
                  <p className="border-b border-black p-2">{process}</p>
                </div>
              );
            })}
          </div>
        </div>
      </div>
    </div>
  );
};
