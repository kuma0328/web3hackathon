import { useRecipeStore } from '../../stores/RecipeStore';
import { TypoGraphy } from '../BaseParts/TypoGraphy';
import { CreateButton } from './CreateButton';
import { DeleteButton } from './DeleteButton';

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
            <CreateButton
              data={material}
              onCreate={setMaterial}
              resetCurrent={() => changeMaterial('')}
            />
          </div>
          <div className="grid grid-cols-3 place-items-center justify-items-center gap-2 pt-3">
            {materials.map((mt, idx) => {
              return (
                <div
                  className="relative w-full rounded-md border border-black p-4 shadow-md"
                  key={idx}
                >
                  <p className="overflow-x-auto">{mt}</p>

                  <DeleteButton onDelete={deleteMaterial} data={mt} />
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
            <CreateButton
              data={process}
              onCreate={setProcess}
              resetCurrent={() => changeProcess('')}
            />
          </div>
          <div>
            {processes.map((pro, idx) => {
              return (
                <div className="flex p-2" key={idx}>
                  <div className="relative mr-10 flex w-10 items-center justify-center rounded-md border border-black p-3 text-center">
                    <DeleteButton onDelete={deleteProcess} data={pro} />
                    {idx}
                  </div>

                  <p className="overflow-x-auto border-b border-black p-2">
                    {pro}
                  </p>
                </div>
              );
            })}
          </div>
        </div>
      </div>
    </div>
  );
};
