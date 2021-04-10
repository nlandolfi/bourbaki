# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make section_measures.pdf
	make graph.pdf

section_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > section_measures.tex

section_measures.pdf: ../../*.tex ../../trademark.pdf *.tex section_measures.tex
	# pdflatex --file-line-error -interaction=nonstopmode section_measures.tex
	../../latexrun section_measures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk_graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: section_measures.pdf
	open section_measures.pdf

o:
	make open

terms: section_measures.tex
	bbk_sheet -mode ts section_measures.tex

spell:
	aspell -c section_measures.tex

s:
	make spell

reset:
	rm -f graph.pdf
	rm -f graph.csv
	rm -f graph.graphviz

clean:
	make reset

remake:
	rm -f section_measures.*
	make reset
	bbk_sheet -mode mk > Makefile