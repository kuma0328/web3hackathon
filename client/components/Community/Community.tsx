import { Board } from '../BaseParts/Board';
type TProps = {
  title: string;
  description: string;
};
export const Community = ({ title, description }: TProps) => {
  return (
    <Board>
      <div className="p-5">
        <div className="m-auto my-3 h-32 w-full rounded-md bg-orange-400"></div>
        <div>
          <p className="text-lg font-semibold">{title}</p>
          <div className="my-5 border-t-2 border-b-2 border-black py-5">
            <p>{description}</p>
          </div>
        </div>
      </div>
    </Board>
  );
};
