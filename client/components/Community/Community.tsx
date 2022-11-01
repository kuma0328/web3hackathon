import { Board } from '../BaseParts/Board';
import { Button } from '../BaseParts/Button';
import { Descriptions } from '../BaseParts/Descriptions';
import { LinkTo } from '../BaseParts/LinkTo';
type TProps = {
  title: string;
  link?: string;
  recipeLink?: string;
  join?: boolean;
  description: string;
  onClickJoin?: () => void;
};
export const Community = ({
  title,
  link,
  recipeLink,
  description,
  onClickJoin,
}: TProps) => {
  return (
    <Board>
      <div className="p-5">
        <div className="m-auto my-3 h-32 w-full rounded-md bg-orange-400"></div>
        <p className="text-center text-lg font-semibold">{title}</p>
        <Descriptions>
          <p>{description}</p>
        </Descriptions>
        <div className="mt-10 flex justify-center gap-5">
          {link !== undefined && (
            <LinkTo link={link} className="m-auto w-2/3 text-center">
              見にいく
            </LinkTo>
          )}
          {recipeLink !== undefined && (
            <LinkTo link={recipeLink} className="m-auto w-2/3 text-center">
              レシピ
            </LinkTo>
          )}
          {onClickJoin !== undefined && (
            <Button onClick={onClickJoin}>参加する</Button>
          )}
        </div>
      </div>
    </Board>
  );
};
