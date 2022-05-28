# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make empty_set.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

empty_set.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > empty_set.tex

empty_set.pdf: ../../*.tex ../../trademark.pdf *.tex empty_set.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode empty_set.tex
	../../latexrun empty_set.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empty_set.pdf
	open empty_set.pdf

o:
	make open

terms: empty_set.tex
	bbk_sheet -mode ts empty_set.tex

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
	rm -f empty_set.*
	make reset
	bbk_sheet -mode mk > Makefile