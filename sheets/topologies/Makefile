# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make topologies.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

topologies.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > topologies.tex

topologies.pdf: ../../*.tex ../../trademark.pdf *.tex topologies.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode topologies.tex
	../../latexrun topologies.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: topologies.pdf
	open topologies.pdf

o:
	make open

terms: topologies.tex
	bbk_sheet -mode ts topologies.tex

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
	rm -f topologies.*
	make reset
	bbk_sheet -mode mk > Makefile