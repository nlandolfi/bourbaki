# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make search_problems.pdf

search_problems.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > search_problems.tex

search_problems.pdf: ../../*.tex ../../trademark.pdf *.tex search_problems.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode search_problems.tex
	../../latexrun search_problems.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: search_problems.pdf
	open search_problems.pdf

o:
	make open

terms: search_problems.tex
	bbk_sheet -mode ts search_problems.tex

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
	rm -f search_problems.*
	make reset
	bbk_sheet -mode mk > Makefile