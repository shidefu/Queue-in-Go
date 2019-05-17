package main

import "errors"

/*
 * Queue
 * kind of data structure, of which the elements are First-In/First-Out,
 * the size can be infinity
 *
 */
type Queue []interface{}

/*
 * push an element into the queue
 *
 * @param element: pushed into the queue
 */
func (queue *Queue) Offer(element interface{}) {
	*queue = append(*queue, element)
}

/*
 * poll an element from the queue
 *
 * @return1: element which was popped
 * @return2: whether the element exists
 *
 * Notes: because the element returned is interface{} type,
 * if other type is required, write the following code in your program,
 * switch v := element.(type) {
 *     case int:
 *         ...
 *     case string:
 *         ...
 *     ...
 * }
 */
func (queue *Queue) Poll() (interface{}, error) {
	if len(*queue) == 0 {
		return nil, errors.New("Queue Over Flow!")
	}
	value := (*queue)[0]
	*queue = (*queue)[1:]
	return value, nil
}

/*
 * get the first element of the queue
 *
 * @return1: first element
 * @return2: whether the element exists
 */
func (queue *Queue) First() (interface{}, error) {
	if len(*queue) == 0 {
		return nil, errors.New("Queue Over Flow!")
	}
	return (*queue)[0], nil
}

/*
 * get the last element of the queue
 *
 * @return1: last element
 * @return2: whether the element exists
 */
func (queue *Queue) Last() (interface{}, error) {
	if len(*queue) == 0 {
		return nil, errors.New("Queue Over Flow!")
	}
	return (*queue)[len(*queue)-1], nil
}

/*
 * judge whether the queue is empty
 */
func (queue *Queue) Size() int {
	return len(*queue)
}
