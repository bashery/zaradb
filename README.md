**Title:** ZaraDB: A Lightweight and Fast Document Database

**Description:**

ZaraDB is a lightweight, simple, and fast document database currently under intensive development. It aims to be a user-friendly alternative to MongoDB, offering a streamlined API for interacting with your data. While official documentation is forthcoming upon stabilization, this README provides a foundational overview.

**Features:**


**examples:**
Here's the converted markdown from the provided HTML:

## Examples

**Insert**

* **Insert one data object:**

```js
{action:"insert", collection:"users", data:{name:"adam", age:12}}
```

* **Insert many data objects (bulk):**

```js
{action:"insertMany", collection:"test", data:[{name:"jalal", age:23},{name:"akram", age:30},{name:"hasna", age:35}]}
```

**Select**

* **Select one object:**

```js
{action:"findOne", collection:"users", match:{name:"adam"}}
```

* **Select objects matching conditions:**

```js
{action:"findMany", collection:"users", match:{name:"adam"}}
```

* **Select objects matching specific conditions (numeric data):**

```js
{action:"findMany", collection:"users", match:{name:"adam", age:{$gt:12}}}

Supported comparison operators: $eq (equal), $nq (not equal), $lt (less than), $gt (greater than), $ge (greater than or equal to), $le (less than or equal to)
```

* **Select objects matching any value:**

```js
{action:"findMany", collection:"users", match:{ age:{$in:[12, 23, 34]}}}
{action:"findMany", collection:"users", match:{ name:{$in:["akram", "zaid"]}}}
```

* **Select objects that do not match any value:**

```js
{action:"findMany", collection:"users", match:{ age:{$nin:[12, 23, 34]}}}
{action:"findMany", collection:"users", match:{ name:{$nin:["akram", "zaid"]}}}
```

* **Select objects matching any conditions (OR operator):**

```js
{action:"findMany", collection:"users", match:{ $or:[name:{$eq:"akram", age:$gt:13}]}}
```

* **Select objects matching all conditions (AND operator):**

```js
{action:"findMany", collection:"users", match:{ $and:[name:{$eq:"akram", age:$gt:13}]}}
```

**Update**

* **Update by ID:**

```js
{action:"updateById", collection:"test", _id:3, data:{name:"hosam", age:10}}
```

* **Update one or more documents matching criteria:**

```js
{action:"updateOne", collection:"test", match:{_id{$gt:33}}, data:{name:"hosam", age:10}}
```

**Delete**

* **Delete the first document matching conditions:**

```js
{action:"deleteOne", collection:"users", match:{name:"adam", age:{$gt:12}}}
```

* **Delete all objects matching conditions:**

```js
{action:"deleteMany", collection:"users", match:{name:"adam", age:{$gt:12}}}
```

* **Skip or ignore some matching objects:**

```js
{action:"deleteMany", collection:"users", match:{name:"adam", age:{$gt:12}}, skip: 3}
```

* **Delete a limited number of matching objects:**

```js
{action:"deleteMany", collection:"users", match:{name:"adam", age:{$gt:12}}, skip: 3, limit:3}
```

* **Exclude fields during retrieval:**

```js
{action:"findMany", collection:"test", fields:{_id:0, name:0}}
```

* **Rename fields during retrieval:**

```js
{action:"findMany", collection:"test", fields:{_id:0, name:"full_name"}}
```

