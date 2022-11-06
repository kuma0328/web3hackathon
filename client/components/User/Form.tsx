import { UserForm } from '../../stores/UserForm';
import { Button } from '../BaseParts/Button';
import { Input } from '../BaseParts/Input';
import { TypoGraphy } from '../BaseParts/TypoGraphy';
import { useLogin } from '../../hooks/User/useLogin';
import { useSignUp } from '../../hooks/User/useSignUp';
import { Meta } from './Meta';
import { useAddress } from '@thirdweb-dev/react';
type TProps = {
  title: string;
};
export const Form = ({ title }: TProps) => {
  const newUser = title === 'ユーザー登録';
  const {
    name,
    mail,
    password,
    setMail,
    setPassword,
    setName,
    passwordConfirmation,
    setPasswordConfirmation,
  } = UserForm();
  const address = useAddress();
  return (
    <div className="m-10">
      <TypoGraphy className="text-center text-xl">{title}</TypoGraphy>
      <Meta address={address ?? ''} />
      {newUser && (
        <Input
          id="name"
          name="名前"
          type="text"
          value={name}
          onChange={setName}
        />
      )}
      <Input
        id="mail"
        name="メールアドレス"
        type="email"
        value={mail}
        onChange={setMail}
      />
      <Input
        id="password"
        name="パスワード"
        type="password"
        value={password}
        onChange={setPassword}
      />
      {newUser && (
        <Input
          id="password-confirmation"
          name="パスワード確認"
          type="password"
          value={passwordConfirmation}
          onChange={setPasswordConfirmation}
        />
      )}

      <div className="my-5 flex justify-center">
        <Button
          onClick={
            newUser
              ? () => useSignUp({ name, mail, password })
              : () => useLogin({ mail, password })
          }
        >
          {title}する
        </Button>
      </div>
    </div>
  );
};
