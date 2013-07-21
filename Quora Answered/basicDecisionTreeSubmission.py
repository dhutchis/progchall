#from sklearn.ensemble import RandomForestClassifier
from sklearn import tree
import json
#from pprint import pprint
import sys
import numpy as np

def form_feature_vec(jso, topicset, topiclist, sample):
    """ Store feature vector in sample from json (not including __ans__) """
    total_topic_followers = 0
    #topic_vec = topiclist_false[:] #start with no-topic vector
    numtopics = 0
    for topic in jso['topics']:
        total_topic_followers += topic['followers']
        numtopics += 1
        #if topic['name'] in topicset:
            #idx = topiclist.index(topic['name'])
            #topic_vec[idx] = True
            #numtopics += 1
    sample[:4] = [
        len(jso['question_text']),
        1 if jso['anonymous'] else 0,
        total_topic_followers,
        numtopics
    ]
    if jso['context_topic'] :#and jso['context_topic']['name'] in topicset:
        sample[4] = topiclist.index(jso['context_topic']['name'])
        sample[5] = jso['context_topic']['followers']
    else:
        sample[4:6] = [-1, 0]
    #sample.extend(topic_vec) #might be faster other way around?
    #return sample


infile = sys.stdin
#infile = open("answered_data_10k.in")
line = infile.readline()
N = int(line[:-1])
train_json = []
# read in N training cases
for i in xrange(0,N):
    s = infile.readline()[:-1]
    train_json.append(json.loads(s))
#pprint(train_json[0], stream=sys.stderr)


line = infile.readline()
T = int(line[:-1])
# read in T test cases
test_json = []
for i in xrange(0,T):
    s = infile.readline()[:-1]
    test_json.append(json.loads(s))
infile.close()


# get set of all topics
topicset = set()
for i in xrange(0,N+T):
    if i < N:
        obj = train_json[i]
    else:
        obj = test_json[i-N]
    ct = obj['context_topic']
    if ct:
        topicset.add(obj['context_topic']['name'])
    for topic in obj['topics']:
        topicset.add(topic['name'])
# 8230 topics in train set, 8771 topics in train+test set
#pprint(len(topicset))
# convert set to list for defined order
topiclist = list(topicset)


# convert json to array of sample (feature) arrays + array of targets
topiclist_false = [False] * len(topiclist)
samples = np.ndarray([N, 6], dtype=int)
targets = np.ndarray([N], dtype=bool)
i = 0
for jso in train_json:
    form_feature_vec(jso, topicset, topiclist, samples[i])
    targets[i] = jso['__ans__']
    i += 1


# Make the classifier!
clf = RandomForestClassifier(n_estimators=50)
clf.fit(samples, targets)


# classify the test data
sample = np.ndarray([6], dtype=int)
for jso in test_json:
    form_feature_vec(jso, topicset, topiclist, sample)
    p = clf.predict(sample)
    #
    #pprint(p, stream=sys.stderr)
    s = '{"__ans__": '
    s += 'true' if p[0] else 'false'
    s += ', "question_key": "' + jso['question_key'] + '"}'
    print(s)
    #jout = json.dumps({"question_key":jso['question_key'],'__ans__':p[0]})
    #print(jout)
    

# version without full topic vector