Quora ML Problem: Answered
==========================
This is a ML competition hosted on [HackerRank](https://www.hackerrank.com/) for [Quora](http://www.quora.com/), hosted on 20 July 2013 between 3:00pm and 11:00pm PST.  I entered with little background knowledge, no setup libraries, and not even an installed Python environment.  For some reason I thought it would be fun to learn the numpy and sklearn Python libraries on the fly and try to implement an ensemble Random Forest algorithm.

I managed to implement everything, extracting some simple features from the training data, training the random forest of decision trees and using them to predict labels on new test data.  Unfortunately, performance was mediocre in both efficiency and outcome.  I could not scale the random forests beyond a few hundred trees for some reason.  Perhaps if the parallel construction worked I would see better results, but HacerRank only allows you to run on a single core (the code executes on HackerRank's servers, not your own).  Ultimately, I had to revert to a single sage decision tree since the random forest construction timed out on HackerRank's servers after 1 minute on their larger data set (unknown size), even with just 10 trees.

The random forests correctly labelled between 57 and 61% of the test data set correctly, a tad better than random guessing.  The final submission with a single decision tree scored 54.94%.

Perhaps after formally studying machine learning this Fall, I will learn how to do a better job.  One idea I could not efficiently implement is to include a sparse binary vector of every unique topic.  There are about 8500 unique topics in the sample data set.




## Problem Description ##
<https://www.hackerrank.com/contests/quora/challenges/quora-ml-answered>
>Quora uses a combination of machine learning algorithms and moderation to ensure high-quality content on the site. High question and answer quality has helped Quora distinguish itself from other Q&A sites on the web.
>
>As we get many questions every day, a challenge we have is to figure out good, interesting and meaningful questions from the bad. What questions are valid ones that can be answered? What questions attract reputable answers that then get upvoted? Can you tell which questions will likely get answers quickly, so that we can surface them in real-time to our users?
>
>For this task, given Quora question text and topic data, predict whether a question gets an upvoted answer within 1 day.
>
>#####Input Format
>
>The first line contains N. N questions follow, each being a valid json object. The following fields of raw data are given in json.
>
>question_key (string): Unique identifier for the question.
>question_text (string): Text of the question.
>context_topic (object): The primary topic of a question, if present. Null otherwise. The topic object will contain a name (string) and followers (integer) count.
>topics (array of objects): All topics on a question, including the primary topic. Each topic object will contain a name (string) and followers (integer) count.
>anonymous (boolean): Whether the question was anonymous.
>__ans__ (boolean): Whether the question got an up-voted answer within 1 day.
>This is immediately followed by an integer T.
>T questions follow, each being a valid json object.
>The json contains all but one field __ans__.
>
>#####Output Format
>
>T rows of JSON encoded fields, with the question_key key containing the unique identifier given in the test data, and the predicted value keyed by __ans__.
>
>#####Constraints
>
>question_key is of ascii format.
>question_text, name in topics and context_topic is of UTF-8 format.
>
> 	0 <= followers <= 106
> 	9000 <= N <= 45000
> 	1000 <= T <= 5000
>
>#####Sample Input
>
> 	9000
> 	json_object
> 	json_object
> 	....
> 	json_object
> 	1000
> 	json_object
> 	json_object
> 	....
> 	json_object
>
>#####Sample Output
>
> 	json_object
> 	json_object
> 	...
> 	json_object
>
>Sample testcases can be downloaded here and used for offline training if desired.
>
>#####Scoring
>
>The answers are evaluated by accuracy.
>
>Number correct classified / Total input size * 100%
>The training and test set each will have approximately an equal number of each boolean type.
>
>Your score will be based only on the hidden input. The sample input is only for your convenience.
>json_object
>1000
>json_object
>json_object
>....
>json_object
>#####Sample Output
>
>json_object
>json_object
>...
>json_object
>Sample testcases can be downloaded here and used for offline training if desired.
>
>#####Scoring
>
>The answers are evaluated by accuracy.
>
>Number correct classified / Total input size * 100%
>The training and test set each will have approximately an equal number of each boolean type.
>
>Your score will be based only on the hidden input. The sample input is only for your convenience.