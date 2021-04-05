const secret_text = "23dn3ir30fd2eddd"

const masking = str => {
  const setLast = str.substr(str.length - 3)
  const setFilter = str.substr(str, str.length - 3).replace(/./g,"*")
  
  return setFilter.concat(setLast)
}

console.log(masking(secret_text))