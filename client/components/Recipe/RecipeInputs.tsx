import { useRecipeStore } from '../../stores/RecipeStore';
import { TypoGraphy } from '../BaseParts/TypoGraphy';

export const RecipeInputs = () => {
  const {
    material,
    process,
    materials,
    processes,
    changeMaterial,
    changeProcess,
    setMaterial,
    setProcess,
    deleteMaterial,
    deleteProcess,
  } = useRecipeStore();
  console.log(processes);
  return (
    <div className="lg:p-5">
      <div className="m-3 grid grid-cols-1 lg:m-10">
        <div className="my-5">
          <div className="flex items-center">
            <TypoGraphy className="my-3 text-lg">材料</TypoGraphy>
            <input
              type="text"
              className="ml-4 rounded-md border border-black py-2 px-5"
              onChange={(e) => changeMaterial(e.target.value)}
              value={material}
            />
            <button
              className="rounded-md border border-black p-2 shadow-lg"
              onClick={() => {
                setMaterial(material);
                changeMaterial('');
              }}
            >
              ＋
            </button>
          </div>
          <div className="grid grid-cols-3 place-items-center justify-items-center gap-2 overflow-y-auto">
            {materials.map((mt) => {
              return (
                <div className="w-full rounded-md border border-black p-4 shadow-md">
                  {mt}
                  <button
                    className="h-6 w-6 rounded-md bg-orange-400"
                    onClick={() => deleteMaterial(mt)}
                  >
                    ×
                  </button>
                </div>
              );
            })}
          </div>
        </div>

        <div className="my-5">
          <div className="flex items-center">
            <TypoGraphy className="my-3 text-lg ">作り方</TypoGraphy>
            <input
              type="text"
              className="ml-4 rounded-md border border-black py-2 px-5"
              onChange={(e) => changeProcess(e.target.value)}
              value={process}
            />
            <button
              className="rounded-md border border-black p-2 shadow-lg"
              onClick={() => {
                setProcess(process);
                changeProcess('');
              }}
            >
              ＋
            </button>
          </div>
          <div className="overflow-y-auto">
            {processes.map((pro, idx) => {
              return (
                <div className="flex p-2">
                  <button
                    className="h-6 w-6 rounded-md bg-orange-400"
                    onClick={() => deleteProcess(pro)}
                  >
                    ×
                  </button>
                  <div className="mr-10 flex w-10 items-center justify-center rounded-md border border-black p-3 text-center">
                    {idx}
                  </div>

                  <p className="border-b border-black p-2">{pro}</p>
                </div>
              );
            })}
          </div>
        </div>
      </div>
    </div>
  );
};
