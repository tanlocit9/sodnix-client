import { useQuery } from '@tanstack/react-query';
import { sleep } from '../../utils/functions';
import { User } from '../../types/user';
export const useUser = (id: number) => {
  return useQuery<User | null>({
    queryKey: ['user', id],
    queryFn: async () => {
      await sleep(1000);
      return {
        name: id === 10 ? 'Phoenix' : 'Sod',
        id,
      };
    },
    enabled: !!id, // Only run if id exists
  });
};