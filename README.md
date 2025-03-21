# poe2-lootfilter

*based on NeverSink's lootfilter v0.1.1 and augmented with a few hundred of my additional rules*

## Basic principles of this project:
1. 💯 <ins>Pick everything you see on the ground, on any area level</ins>. At AreaLevel *1+ / 65+ / 71+ / 75+* you see fewer and fewer items that have str- & dex-requirements on them and other useless items for your class. <br> **Characters of level 90+**: check out the "Endgame custom rules" section to apply additional rules for hiding potentially excessive items.
2. 🖋️ <ins>Minimalism for equipment styles</ins>: ~99% of the equipment is garbage that doesn't deserve much custom styling, but you'll see a colorful 🎨 design of currency, map-specific items etc.
Special equipment stylization applies in four cases:
   - for the best bases of the corresponding slot, on which the maximum affix tier *(itemLevel 81–82)* can be generated;
   - for some infrequent items on low-level characters: jewelry, belts, and Chiming staff;
   - items with additional quality *(caster weapons and armors only)*;
   - bases for crafting uniques with Chance Orb.
3. 📖 <ins>Ultimate documentation</ins>: in just a couple minutes you can familiarize yourself with the styles of ~95% of the items in the game and ~95% of the rules for showing/hiding items - there are over a dozen screenshots of how lootfilter works at the end of the article. If you're only interested in individual rules/styles, you can find them in the table of contents (look at the beginning of the lootfilter file)


More details are listed inside the `.lootfilter`-file.<br>
For now, I've only created one version of lootfilter — 🧙‍♀️ **for caster classes**.<br>
*Note*: is not for a specific caster class, but **for the gameplay style** of your chosen class 🧔🧝🐻

## INSTALLATION / UPDATE:

### ☁️ CLOUD FILE
*Recommended for characters at levels 1-90.*<br>
Automatically updated lootfilter without the need to download any files.

🧾 **Instruction:**
1. Log in to the official website: https://www.pathofexile.com/
2. Open the page with list of lootfilters from the author of this repository: [fpsthirty-6961/item-filters](https://www.pathofexile.com/account/view-profile/fpsthirty-6961/item-filters)
3. Click on the "Follow" button in lootfilter you're interested in. Now there are two variants of lootfilters:
   - for casters (wands and focuses are always shown)
   - for summoners (scepters and int-shields are always shown)
4. Turn on filter in the game settings : `Escape` -> Options -> Game -> select the filter from the Dropdown box of "Item Filter" section — name of lootfilter will be **blue** color.

  ![settings/game-itemfilter](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/settings/game-itemfilter.png)


### 🗄️ LOCAL FILE
*Recommended for characters at levels 90+.*<br>
In this case you can make your own edits, including activating additionally prepared rules in the 📋 "[[0700]] Endgame custom rules" section.

🧾 **Instruction:**
1. Open link of current release: [releases/latest](https://github.com/fpsthirty/poe2-lootfilter/releases/latest)
2. Scroll down to the "Assets"-section and click on the file name `NvrSnk+fps30_caster.filter` to download it.
3. Paste this file into the following folder: `%userprofile%/Documents/My Games/Path of Exile 2` *(you can paste this correct link into File Explorer)*
4. Turn on filter in the game settings : `Escape` -> Options -> Game -> select the filter from the Dropdown box of "Item Filter" section — name of lootfilter will be **white** color.

> [!WARNING]  
> Summoner settings are enabled by default; if you are playing as a caster without a large number of summons, open lootfilter-file in any text editor (I recommend Sublime Text), find "[[0200]] Section for non-summoner casters" section outside the table of contents.

## 📖 Documentation
*click on the last dropdown item first, you'll be more comfortable, trust me.*

<details>
  <summary>show/hide rules for armors in the different area levels</summary>
> [!NOTE] 
> - double values in the lines for Shields indicate the difference in display depending on selected profile: summoner / caster;
> - some item bases you've never seen, but information about them is in the game.
> [!NOTE]  
> An example of how to navigate the table:
> there is an armor item that has "Greathelm" in its name, and depending on the rarity of the item, different rules are applied to show/hide it:
> ⚪ normal Greathelm: will be hidden at any area level (1-99);
> 🔵 magic Greathelm: will only show on area levels 1-70;
> 🟡 rare Greathelm: will only show on area levels 1-74;
> 🟠 unique Greathelm: will be shown in any area level (1-99).
<br>
<br>![table-armors](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/table-armors.png)
</details>
<details>
  <summary>show/hide rules for weapons in the different area levels</summary>
> [!NOTE]  
> - double values in the lines for Focuses, Sceptres and Wands indicate the difference in display depending on selected profile: summoner / caster;
> - some item bases you've never seen, but information about them is in the game.
> [!TIP]
> * the Chiming Staff is always displayed
> [!NOTE]  
> An example of how to navigate the table:
<br> An example of how to navigate the table:
<br> there is an weapon item that has "Crossbow" in its name, and depending on the rarity of the item, different rules are applied to show/hide it:
<br> ⚪ normal Crossbow: will be hidden at any area level (1-99);
<br> 🔵 magic Crossbow: will only show on area levels 1-64;
<br> 🟡 rare Crossbow: will only show on area levels 1-74;
<br> 🟠 unique Crossbow: will be shown in any area level (1-99).
<br>
<br>![table-weapons](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/table-weapons.png)
</details>
<details>
  <summary>altar</summary>
  ![altar](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/altar.png)
</details>
<details>
  <summary>breach</summary>
  ![breach](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/breach.png)
</details>
<details>
  <summary>currency</summary>
  ![currency](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/currency.png)
</details>
<details>
  <summary>delirium</summary>
  ![delirium](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/delirium.png)
</details>
<details>
  <summary>equipment: high-end</summary>
  -end![equip_high-end](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/equip_high-end.png)
</details>
<details>
  <summary>equipment</summary>
  ![equipment](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/equipment.png)
</details>
<details>
  <summary>essences</summary>
  ![essences](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/essences.png)
</details>
<details>
  <summary>expedition</summary>
  ![expedition](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/expedition.png)
</details>
<details>
  <summary>gold</summary>
  ![gold](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/gold.png)
</details>
<details>
  <summary>pinnacle keys</summary>
  ![pinnacle_keys](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/pinnacle_keys.png)
</details>
<details>
  <summary>socketables</summary>
  ![socketables](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/socketables.png)
</details>
<details>
  <summary>tablets</summary>
  ![tablets](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/tablets.png)
</details>
<details>
  <summary>trials overall</summary>
  ![trials_overall](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/trials_overall.png)
</details>
<details>
  <summary>uncut_gems</summary>
  ![uncut_gems](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/image.png)
</details>
<details open>
  <summary>waystones | next your click to one dropdown higher, etc. — so you have to scroll less. </summary>
  ![waystones](https://github.com/fpsthirty/poe2-lootfilter/raw/master/img/loot/image.png)
</details>
<script>
document.addEventListener("DOMContentLoaded", () => {
    document.querySelectorAll("details").forEach((details) => {
        details.addEventListener("toggle", function () {
            if (this.open) {
                document.querySelectorAll("details").forEach((otherDetails) => {
                    if (otherDetails !== this) {
                        otherDetails.open = false;
                    }
                });
            }
        });
    });
});
</script>

---

> [!NOTE]  
💬 I express my respect to **@NeverSink** for the **[original version](https://github.com/NeverSinkDev/NeverSink-Filter-for-PoE2)** of lootfilter and will mention its credentials below:
* TWITTER: @NeverSinkDev
* DISCORD: https://discord.gg/mye6xhF
* TWITCH:  https://www.twitch.tv/neversink
