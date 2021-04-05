// fetch is not defined if run in server
// run in the browser (not server), open file index.html in browser and open console

const posts = fetch('https://jsonplaceholder.typicode.com/posts')
  .then((response) => response.json())
  .then((data) => data);

const users = fetch('https://jsonplaceholder.typicode.com/users')
  .then((response) => response.json())
  .then((data) => data);

const customApi = async (post, user) => {
  const dataUsers = await user;
  const dataPosts = await post;

  const result = await dataPosts.map(item1 => ({...item1, user: dataUsers.find(item2 => item2.id === item1.userId)}))

  console.log(result, ' ... result')
}

customApi(posts, users);