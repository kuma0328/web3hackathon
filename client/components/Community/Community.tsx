import { Button } from '../BaseParts/Button';
import { Descriptions } from '../BaseParts/Descriptions';
import { LinkTo } from '../BaseParts/LinkTo';
import Image from 'next/image';
type TProps = {
  title: string;
  link?: string;
  recipeLink?: string;
  join?: boolean;
  description: string;
  onClickJoin?: () => void;
  image: string;
  className?: string;
};
export const Community = ({
  title,
  link,
  recipeLink,
  description,
  onClickJoin,
  className,
  image,
}: TProps) => {
  console.log(recipeLink);
  // Todo リンクが動かない
  return (
    <div className="p-5">
      <div className={`${className}`}>
        <Image src={image} width={500} height={500} alt="イメージ画像" />
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
    </div>
  );
};
