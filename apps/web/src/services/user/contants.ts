export const USER_KEYS = {
  all: ['users'] as const,
  lists: () => [...USER_KEYS.all, 'list'] as const,
  list: (page: number, limit: number) => [...USER_KEYS.lists(), page, limit] as const,
  details: () => [...USER_KEYS.all, 'detail'] as const,
  detail: (id: number) => [...USER_KEYS.details(), id] as const,
};
