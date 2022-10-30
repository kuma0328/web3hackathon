import { Form } from '../../components/User/Form';

export const signup = () => {
  return (
    <div className="flex justify-center">
      <Form
        mail={''}
        password={''}
        onChangeMail={() => {
          console.log('aa');
        }}
        onChangePassword={() => {
          console.log('aa');
        }}
        onSubmit={() => {
          console.log('aa');
        }}
      />
    </div>
  );
};
export default signup;
