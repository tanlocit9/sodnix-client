import { ROUTES } from '../../contants/route';
import { useLogic } from './useLogic';
import { twMerge } from 'tailwind-merge';

export const Sidebar = () => {
  const { isActive, handleClick } = useLogic();
  return (
    <aside className="w-64 bg-gray-200 p-4 shadow-md">
      <ul className="list-none p-0">
        {ROUTES &&
          Object.entries(ROUTES).map(([key, route]) => {
            return (
              <li key={route.path} className="mb-2">
                <div
                  className={twMerge(
                    'flex items-center px-4 py-2 rounded-md hover:bg-gray-300 dark:hover:bg-gray-700 cursor-pointer',
                    isActive(route.path) &&
                      'bg-gray-300 dark:bg-gray-700 text-gray-900 dark:text-white font-semibold',
                  )}
                  onClick={() => handleClick(route.path)}
                >
                  <span className="mx-4 font-medium">{route.name}</span>
                </div>
              </li>
            );
          })}
      </ul>
    </aside>
  );
};
