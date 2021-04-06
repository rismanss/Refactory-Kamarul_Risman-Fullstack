// fetch is not defined if run in server
// run in the browser (not server), open file index.html in browser and open console

const fetchAPI = ({ endpoint }) => {
  return fetch(`https://jsonplaceholder.typicode.com/${endpoint}`)
  .then((response) => response.json())
  .then((data) => data)
  .catch(err => console.error(err));
}

const customApi = () => {
  const posts = fetchAPI({ endpoint: 'posts'});
  const users = fetchAPI({ endpoint: 'users'});

  return Promise.all([posts, users]).then(val => {
    const [postsData, usersData] = val;
    return postsData.map(item1 => ({...item1, user: usersData.find(item2 => item2.id === item1.userId)}));
  }).catch(err => {
    console.error(err);
  });
}

customApi().then(res => console.log(res));