import { useRouter } from 'next/router';
const communityDetail = () => {
  const router = useRouter();
  const { id } = router.query;
  return <div>communityDetail</div>;
};

export default communityDetail;
