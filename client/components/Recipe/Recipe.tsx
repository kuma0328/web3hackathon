import { TypoGraphy } from '../BaseParts/TypoGraphy';

export const Recipe = () => {
  const contents = [
    '醤油',
    '鶏肉',
    '牛肉',
    'オリーブオイル',
    'ニンニクを炒めたのもの',
    '何か',
    '何か',
    '何か',
    '何か',
    '何か',
    '何か',
    '何か',
  ];
  const processes = [
    'ニンニクを5分炒める',
    '肉を弱火で炒める',
    'オリーブオイルをかけて蒸す',
    '肉とニンニクを混ぜ合わせる',
    '醤油で味付けする',
    '盛り付けする',
    '盛り付けする',
    '盛り付けする',
    '盛り付けする',
  ];
  return (
    <div className="lg:p-5">
      <div>
        <TypoGraphy className="my-3 text-lg">材料</TypoGraphy>
      </div>
      <div className="m-3 grid grid-cols-1 lg:m-10">
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
