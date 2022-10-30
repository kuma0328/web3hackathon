import { Board } from '../BaseParts/Board';
import { Button } from '../BaseParts/Button';
import { Descriptions } from '../BaseParts/Descriptions';
import { LinkTo } from '../BaseParts/LinkTo';
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
          <p className="text-center text-lg font-semibold">{title}</p>
          <Descriptions>
            <p>{description}</p>
          </Descriptions>
        </div>
        <LinkTo link="aaa" className="m-auto w-2/3 text-center">
          見にいく
        </LinkTo>
      </div>
    </Board>
  );
};