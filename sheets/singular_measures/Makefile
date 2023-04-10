# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make singular_measures.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

singular_measures.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > singular_measures.tex

singular_measures.pdf: ../../*.tex ../../trademark.pdf *.tex singular_measures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode singular_measures.tex
	../../latexrun singular_measures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: singular_measures.pdf
	open singular_measures.pdf

o:
	make open

terms: singular_measures.tex
	bbk_sheet -mode ts singular_measures.tex

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
	rm -f singular_measures.*
	make reset
	bbk_sheet -mode mk > Makefile