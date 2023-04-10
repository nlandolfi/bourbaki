# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make empirical_normal.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

empirical_normal.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > empirical_normal.tex

empirical_normal.pdf: ../../*.tex ../../trademark.pdf *.tex empirical_normal.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode empirical_normal.tex
	../../latexrun empirical_normal.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: empirical_normal.pdf
	open empirical_normal.pdf

o:
	make open

terms: empirical_normal.tex
	bbk_sheet -mode ts empirical_normal.tex

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
	rm -f empirical_normal.*
	make reset
	bbk_sheet -mode mk > Makefile