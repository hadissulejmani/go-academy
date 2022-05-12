
const host = 'http://localhost:3000';
const prefix = 'api';

export const baseUrl = [host, prefix].join('/');

const routes = {
  lists: () => ['lists'].join('/'),
  tasks: () => ['tasks'].join('/'),
  listTasks: (id) => ['lists', id, 'tasks'].join('/'),
  list: (id) => ['lists', id].join('/'),
  task: (id) => ['tasks', id].join('/'),
};

export default routes;
