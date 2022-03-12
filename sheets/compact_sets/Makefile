# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make compact_sets.pdf

compact_sets.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > compact_sets.tex

compact_sets.pdf: ../../*.tex ../../trademark.pdf *.tex compact_sets.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode compact_sets.tex
	../../latexrun compact_sets.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: compact_sets.pdf
	open compact_sets.pdf

o:
	make open

terms: compact_sets.tex
	bbk_sheet -mode ts compact_sets.tex

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
	rm -f compact_sets.*
	make reset
	bbk_sheet -mode mk > Makefile