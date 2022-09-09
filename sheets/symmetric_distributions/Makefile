# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make symmetric_distributions.pdf

symmetric_distributions.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > symmetric_distributions.tex

symmetric_distributions.pdf: ../../*.tex ../../trademark.pdf *.tex symmetric_distributions.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode symmetric_distributions.tex
	../../latexrun symmetric_distributions.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: symmetric_distributions.pdf
	open symmetric_distributions.pdf

o:
	make open

terms: symmetric_distributions.tex
	bbk_sheet -mode ts symmetric_distributions.tex

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
	rm -f symmetric_distributions.*
	make reset
	bbk_sheet -mode mk > Makefile