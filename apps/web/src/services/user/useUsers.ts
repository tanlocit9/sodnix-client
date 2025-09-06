import { USER_KEYS } from './contants';
import { useQuery, useMutation, useQueryClient } from '@tanstack/react-query';
import { sleep } from '../../utils/functions';
export const useUsers = (page = 1, limit = 10) => {
  return useQuery({
    queryKey: USER_KEYS.list(page, limit),
    queryFn: async () => {
        await sleep(500);
        return [];
    },
  });
};