# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make metrics.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

metrics.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > metrics.tex

metrics.pdf: ../../*.tex ../../trademark.pdf *.tex metrics.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode metrics.tex
	../../latexrun metrics.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: metrics.pdf
	open metrics.pdf

o:
	make open

terms: metrics.tex
	bbk_sheet -mode ts metrics.tex

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
	rm -f metrics.*
	make reset
	bbk_sheet -mode mk > Makefile