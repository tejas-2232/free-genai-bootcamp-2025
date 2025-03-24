# SQLite3 Best Practices for Junior Developers

This guide outlines the three most important rules to follow when working with SQLite3 in Python applications.

## 1. Always Use Parameterized Queries

Never construct SQL queries by string concatenation. Always use parameterized queries with placeholders to prevent SQL injection attacks.

```python
# ✅ GOOD: Using parameterized queries
cursor.execute('SELECT * FROM words WHERE id = ?', (word_id,))

# ❌ BAD: String concatenation (SQL injection risk!)
cursor.execute(f'SELECT * FROM words WHERE id = {word_id}')
```

### Why?
- Prevents SQL injection attacks
- Handles data type conversion automatically
- Improves query performance through query plan caching
- Makes code more maintainable

## 2. Proper Connection and Transaction Management

Always manage database connections and transactions properly to prevent resource leaks and ensure data consistency.

```python
# ✅ GOOD: Using context manager
with sqlite3.connect('database.db') as connection:
    cursor = connection.cursor()
    cursor.execute('INSERT INTO words (kanji, romaji) VALUES (?, ?)', (kanji, romaji))
    connection.commit()

# ✅ GOOD: Explicit connection management
try:
    connection = sqlite3.connect('database.db')
    cursor = connection.cursor()
    cursor.execute('INSERT INTO words (kanji, romaji) VALUES (?, ?)', (kanji, romaji))
    connection.commit()
finally:
    connection.close()

# ❌ BAD: No connection management
connection = sqlite3.connect('database.db')
cursor = connection.cursor()
cursor.execute('INSERT INTO words (kanji, romaji) VALUES (?, ?)', (kanji, romaji))
# Connection never closed!
```

### Why?
- Prevents connection leaks
- Ensures transactions are properly committed
- Handles errors gracefully
- Maintains data consistency

## 3. Handle Foreign Key Constraints Properly

Always define and respect foreign key relationships in your schema. When deleting records, handle dependent records first.

```python
# ✅ GOOD: Proper handling of foreign key constraints
try:
    # First delete child records
    cursor.execute('DELETE FROM word_review_items')
    # Then delete parent records
    cursor.execute('DELETE FROM study_sessions')
    connection.commit()
except sqlite3.IntegrityError as e:
    connection.rollback()
    raise e

# ❌ BAD: Ignoring foreign key constraints
cursor.execute('DELETE FROM study_sessions')  # Will fail if there are dependent records
```

### Why?
- Maintains referential integrity
- Prevents orphaned records
- Makes database maintenance easier
- Ensures data consistency

## Additional Tips

1. **Use Row Factory**
   ```python
   connection.row_factory = sqlite3.Row  # Returns rows as dictionaries
   ```

2. **Enable Foreign Keys**
   ```python
   connection.execute('PRAGMA foreign_keys = ON')
   ```

3. **Use Transactions for Multiple Operations**
   ```python
   try:
       connection.execute('BEGIN TRANSACTION')
       # Multiple operations here
       connection.commit()
   except:
       connection.rollback()
       raise
   ```

4. **Index Your Tables**
   ```sql
   CREATE INDEX idx_words_kanji ON words(kanji);
   ```

Remember: Following these practices will help you build more secure, reliable, and maintainable applications with SQLite3.