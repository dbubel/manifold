package mocks

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	proto "github.com/dbubel/manifold/proto_files"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProduceCommand struct {
	Args []string
}

func (c *ProduceCommand) Help() string {
	return ""
}

func (c *ProduceCommand) Synopsis() string {
	return ""
}

func (c *ProduceCommand) Run(args []string) int {
	conn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials())) // grpc.WithBlock()
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return 0
	}

	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	x := proto.NewManifoldClient(conn)
	ts := time.Now()

	totalMsg, _ := strconv.Atoi(c.Args[1])
	concurrency, _ := strconv.Atoi(c.Args[2])
	// fmt.Println(lcm(totalMsg, concurrency), concurrency, lcm(totalMsg, concurrency)/concurrency)

	var wg sync.WaitGroup
	totalMsg = (totalMsg / concurrency) * concurrency
	wg.Add((totalMsg / concurrency) * concurrency)

	for i := 0; i < concurrency; i++ {
		go func() {
			for i := 0; i < totalMsg/concurrency; i++ {
				fmt.Println("sending")
				time.Sleep(time.Millisecond * 500)
				_, err := x.Enqueue(context.Background(), &proto.EnqueueMsg{
					Priority:  proto.Priority_NORMAL,
					TopicName: "hello23",
					// Data:      []byte(fmt.Sprintf("{\totalMsg\t\"postId\": \"postId-12534563344664\",\totalMsg\t\"html\": \"here’s more to Maine than lobsters and Stephen King novels. It’s also home to some of the best hiking in the northeast US, with thousands of miles of trails.\\r\\totalMsg\\r\\nBut choosing the best hiking routes in Maine is no easy feat. For the adventurous explorer, there are plenty of backcountry multi-day hikes and for the beginners who just want to see the beautiful coastline, the Pine Tree State has you covered. \\r\\totalMsg\\r\\nFrom coastal treks with epic views of lighthouses perched over rocky saltwater-drenched ledges to strenuous bucket list hikes in the state's legendary wilderness, lace up your hiking boots and discover why Maine is called Vacationland.\\r\\totalMsg\\r\\nExplore the planet's most surprising adventures with our weekly newsletter delivered to your inbox.\\r\\n1. Beehive Trail, Acadia National Park\\r\\nBest for a unique hiking experience \\r\\n1.5 miles, 2–3 hours, strenuous\\r\\totalMsg\\r\\nAcadia National Park offers no shortage of incredible mountain peaks to climb for breathtaking views of the Atlantic Ocean and Maine’s rocky coastline. The Beehive Trail is one of the most unique trails in Acadia National Park and much of the trail requires climbing up iron rung ladders bolted to the granite.\\r\\totalMsg\\r\\nThe trailhead begins on a gradual path through the forest until you come to a trail marker. You’ll begin to climb some granite stairs, over iron bridges and finally up a series of iron rung ladders over boulders. Once you reach the summit at 520ft, you’re rewarded with stunning views of Sand Beach. You’ll descend the Bowl Trail where you can cool down in the Bowl, an alpine pond between Beehive and Gorham Mountain.\\r\\totalMsg\\r\\n2. Mount Katahdin, Baxter State Park\\r\\nBest for experienced hikers  \\r\\n5.2 miles, 8–12 hours, very strenuous\\r\\totalMsg\\r\\nNamed the “Greatest Mountain” by the Penobscot, Mt Katahdin is located in the heart of Baxter State Park. Standing at 5269ft, this is both Maine’s tallest peak and the northern terminus of the Appalachian Trail. Hiking Mt Katahdin takes approximately 8 to 12 hours and is a very strenuous hike.\\r\\totalMsg\\r\\nThe 5.2-mile Hunt Trail is one of the most popular trails to reach Baxter Peak as it offers picturesque views of Maine’s foothills and Katahdin Stream Falls. For well-prepared, experienced hikers, the famous 1.1-mile transverse of Knife Edge will test your fear of heights, but you’ll earn bragging rights with the locals.\\r\\totalMsg\\r\\nWoman smiling as she jumps into water with a waterfall thundering in the background and a man laughing as he watches her plunge\\r\\nBreak up your hike with a swim at Gulf Hagas © Chris Bennett \\/ Getty Images\\r\\n3. Gulf Hagas, Brownville\\r\\nBest hike for going for a swim\\r\\n8.2 miles, 5–6 hours, moderate\\r\\totalMsg\\r\\nNicknamed the “Grand Canyon of the East,” Gulf Hagas on the West Branch of the Pleasant River is a three-mile-long rock canyon that towers 500ft above the bubbling river. Gulf Hagas can be accessed through the Katahdin Iron Works Road in Brownville as part of the Appalachian Trail corridor in central Maine.\\r\\totalMsg\\r\\nShortly after the parking area, you’ll need to ford the river. The water levels vary widely depending on the season and rainfall. From here, you’ll walk along the Appalachian Trail through 150-year-old white pines in the Hermitage before connecting to the Gulf Hagas loop. At the next trail intersection, choose the Rim Trail, so you’re facing the numerous waterfalls as you ascend the loop trail. On a hot summer's day there are plenty of opportunities to go for a swim in the cool water.\\r\\totalMsg\\r\\n4. Mt Battie, Camden Hills State Park\\r\\nBest for the family \\r\\n1.1 miles, 1–2 hours, moderate \\r\\totalMsg\\r\\nMt Battie, in Camden Hills State Park, is one of Maine’s most iconic hikes. The 1.1-mile trail is short and steep and requires some scrambling but is doable for people of all ages. Standing at 780ft, the summit of Mt Battie features stunning views of the town of Camden and Penobscot Bay. Climb the stone tower for even better views and selfies.\\r\\totalMsg\\r\\n5. Fairy Head Loop, Cutler Coast Public Reserved Land\\r\\nBest for nature enthusiasts \\r\\n10.4 miles, 7–8 hours, difficult\\r\\totalMsg\\r\\nOverlooking the Bay of Fundy, Cutler Coast Public Reserved Land comprises over 12,334 acres of wilderness in Downeast Maine. Known as the Bold Coast, the nature preserve is a unique place to experience a variety of Maine’s coastal ecosystems. Bring your camera as you don’t want to miss the sunrise over the rocky coastline.\\r\\totalMsg\\r\\nThe Fairy Head Loop is a 10.4-mile loop trail that provides almost four miles of shorefront hiking before running inland through meadows, forests and grass marshes. Wildlife is abundant, and there are a few camping spots available on a first-come, first-served basis.\\r\\totalMsg\\r\\nA man stands with his small dog at the summit of a mountain looking out at the view of a distant lake\\r\\nThe moderate climb up Tumbledown is a popular day hike © Tennyson Tappan \\/ Getty Images\\r\\n6. Tumbledown Mountain, Weld\\r\\nBest hike for joining the crowd  \\r\\n3.7 miles, 3–4 hours, moderate\\r\\totalMsg\\r\\nNestled among the forested peaks in the western mountains of Maine, Tumbledown Mountain in Weld is one of Maine’s most popular day hikes. This isn't the highest peak in the area or the peak with the best views, but it does have a couple of real showstoppers – the alpine pond situated at 2800ft and the 700ft granite cliffs on the south face overlooking the pond.\\r\\totalMsg\\r\\nThe easiest and most direct trail to the pond is the Brook Trail, which features a 1600ft elevation gain. Tumbledown Ridge Trail transverses towards the East Peak and descends a saddle until it climbs to the summit of West Peak.\\r\\totalMsg\\r\\nThe Loop Trail is the most difficult and recommended for experienced hikers as it's very steep in sections, and you’ll need to climb up metal rungs through boulders to reach the summit. Bring your swimsuit and a picnic, and enjoy the peace from the granite summit.\\r\\totalMsg\\r\\n7. 100-Mile Wilderness, Central Maine\\r\\nBest for experienced multi-day hikers \\r\\n100 miles, 5–10 days, strenuous\\r\\totalMsg\\r\\nSpanning 100 miles from the small central Maine town of Monson to the southern border of Baxter State Park, the 100-Mile Wilderness is often called the “wildest section” of the Appalachian Trail as it’s both challenging to navigate and to traverse. Best to hike late June through July, the 100-Mile Wilderness is a true bucket-list hike for experienced and adventurous hikers.\\r\\totalMsg\\r\\nYou’ll need to pack everything you need and you should expect to be trekking for 8 to 12 hours a day. Throughout the 100 miles, you’ll climb almost 15,000ft. While the hike can be grueling, it is an incredible opportunity to experience Maine’s wilderness and wildlife. Keep your eyes open for moose.\\r\\totalMsg\\r\\nSunset from the Appalachian Trail, Bigelow Mountain, Maine\\r\\nBigelow Mountain on the Appalachian Trail is ideal for experienced hikers © Cavan Images \\/ Getty Images\\r\\n8. Bigelow Mountain, Bigelow Preserve\\r\\nBest for elite hikers \\r\\n16.3 miles, 8–10 hours, strenuous\\r\\totalMsg\\r\\nHead off the beaten path and experience one of Backpacker Magazine’s hardest day hikes in America with the Bigelow Mountain Traverse. The 16.3-mile traverse of the Bigelow Mountain Range via the Appalachian Trail offers some of the most incredible views of Maine’s western mountains and nearby Flagstaff Lake.\\r\\totalMsg\\r\\nThe quad-burning hike is a point-to-point hike, so you’ll need to plan accordingly with cars at two points. The ascent of Little Bigelow begins gradually before turning steeper until you reach the summit of Little Bigelow. From here, the next 6.4 miles are grueling ascends and descends along Bigelow Mountain to Avery Peak at 4088ft. There are more ups and downs until the end. If you’d like to turn the day hike into a weekend hike, there are plenty of tent platforms just below Avery Peak in Bigelow Col. \\r\\totalMsg\\r\\n9. Southwest Ridge Trail, Pleasant Mountain\\r\\nBest for stunning views \\r\\n5.8 miles, 3–4 hours, moderate \\r\\totalMsg\\r\\nJust an hour’s drive from Portland, Pleasant Mountain is southern Maine’s tallest mountain standing at 2006ft. Managed by Loon Echo Land Trust, Pleasant Mountain is home to six trails covering over 10 miles. The views of the open summit are abundant, and if you’re lucky, you may be able to spot Mt Washington in New Hampshire. \\r\\totalMsg\\r\\nThe 3.6-mile Ledges Trail is the most popular and direct trail to the summit, but on a beautiful summer's day, it will be packed with fellow hikers. If you prefer the solitude of nature, the 5.8-mile roundtrip Southwest Ridge Trail is a great choice. Pack a picnic lunch and relax on the granite ledges of the summit.\\r\\totalMsg\\r\\n10. Cadillac Mountain, Acadia National Park\\r\\nBest for watching the sunrise \\r\\n2.2 miles, 2–4 hours, moderate\\r\\totalMsg\\r\\nCadillac Mountain in Acadia National Park is one of the first points in the US to see the sunrise. While most people drive to the summit, the hike is relatively easy compared to other hikes on Mount Desert Island. The open granite peak of Cadillac Mountain offers almost panoramic views of Bar Harbor, Mount Desert Island and the Atlantic Ocean.\\r\\totalMsg\\r\\nThe best time to hike here is in the middle of the night so that you are at the summit and ready to see the sunrise over America. Grab a headlamp and head out along the 2.2-mile Cadillac Mountain North Ridge Trail to the 1528-ft summit under the stars. Don’t forget your flask of coffee! \",\totalMsg\t\"title\": \"The top 10 hiking trails in Maine\",\totalMsg\t\"url\": \"https://www.smoothbrain123ss453.com\",\totalMsg\t\"workspace\": \"content-testing\",\totalMsg\t\"writeKey\": \"{{ _.writeKey }}\"\totalMsg} %d", i)),
					Data: []byte(fmt.Sprintf("hello %d", i)),
				})
				if err != nil {
					log.Fatalf("%v.MyStreamingMethod(_) = _, %v", c, err)
				}
				wg.Done()
			}
		}()
	}

	wg.Wait()
	LEN, _ := x.TopicLength(context.Background(), &proto.DequeueMsg{
		TopicName: "hello23",
	})
	t := time.Now().Sub(ts).Milliseconds()
	msPerMsg := float64(t) / float64(totalMsg)
	// fmt.Println("msPerMsg,totalTime,totalMsg,concurrency,lenQueueServer")
	fmt.Printf("%f,%d,%d,%d,%d\n", msPerMsg, t, totalMsg, concurrency, LEN.Length)
	// fmt.Println(float64(time.Now().Sub(ts).Milliseconds())/float64(totalMsg), time.Now().Sub(ts))
	// fmt.Println("FINAL LEN", LEN)
	// x.DeleteTopic(context.Background(), &proto.DeleteTopicMsg{TopicName: "hello23"})
	return 1
}
