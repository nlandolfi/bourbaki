# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make box_topologies.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

box_topologies.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > box_topologies.tex

box_topologies.pdf: ../../*.tex ../../trademark.pdf *.tex box_topologies.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode box_topologies.tex
	../../latexrun box_topologies.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: box_topologies.pdf
	open box_topologies.pdf

o:
	make open

terms: box_topologies.tex
	bbk_sheet -mode ts box_topologies.tex

spell:
	aspell -c sheet.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f box_topologies.*
	make reset
	bbk_sheet -mode mk > Makefile