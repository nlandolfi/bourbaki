# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make monotone_neighborhoods.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

monotone_neighborhoods.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > monotone_neighborhoods.tex

monotone_neighborhoods.pdf: ../../*.tex ../../trademark.pdf *.tex monotone_neighborhoods.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode monotone_neighborhoods.tex
	../../latexrun monotone_neighborhoods.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: monotone_neighborhoods.pdf
	open monotone_neighborhoods.pdf

o:
	make open

terms: monotone_neighborhoods.tex
	bbk_sheet -mode ts monotone_neighborhoods.tex

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
	rm -f monotone_neighborhoods.*
	make reset
	bbk_sheet -mode mk > Makefile