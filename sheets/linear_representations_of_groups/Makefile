# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make linear_representations_of_groups.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

macros.tex: sheet.lit
	bbk macros sheet.lit > macros.tex

linear_representations_of_groups.tex: sheet.tex macros.tex
	bbk_sheet -mode c -in sheet.tex > linear_representations_of_groups.tex

linear_representations_of_groups.pdf: ../../*.tex ../../trademark.pdf *.tex linear_representations_of_groups.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode linear_representations_of_groups.tex
	../../latexrun linear_representations_of_groups.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: linear_representations_of_groups.pdf
	open linear_representations_of_groups.pdf

o:
	make open

terms: linear_representations_of_groups.tex
	bbk_sheet -mode ts linear_representations_of_groups.tex

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
	rm -f linear_representations_of_groups.*
	make reset
	bbk_sheet -mode mk > Makefile