import { useLocation, useNavigate } from "react-router-dom";

export const useLogic = () => {
  const location = useLocation();
	const navigate = useNavigate();

  const isActive = (path: string) => location.pathname === path;

	const handleClick = (path: string) => {
		navigate(path);
	};

  return {
    isActive,
    handleClick,
  };
};
