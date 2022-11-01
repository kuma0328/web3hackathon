import { Board } from '../BaseParts/Board';
import { Descriptions } from '../BaseParts/Descriptions';
import { LinkTo } from '../BaseParts/LinkTo';
type TProps = {
  title: string;
  link?: string;
  description: string;
};
export const Community = ({ title, link, description }: TProps) => {
  return (
    <Board className="m-2">
      <div className="p-5">
        <div className="m-auto my-3 h-32 w-full rounded-md bg-orange-400"></div>
        <div>
          <p className="text-center text-lg font-semibold">{title}</p>
          <Descriptions>
            <p>{description}</p>
          </Descriptions>
        </div>
        {link !== undefined && (
          <LinkTo link={link} className="m-auto w-2/3 text-center">
            見にいく
          </LinkTo>
        )}
      </div>
    </Board>
  );
};
