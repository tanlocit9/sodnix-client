import { useUser } from '../../services/user/useUser';
export const useLogic = () => {
  const {data: user, isLoading} = useUser(10);
  // Your logic here
  return { user, isLoading };
};
