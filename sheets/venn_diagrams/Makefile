# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make venn_diagrams.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

venn_diagrams.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > venn_diagrams.tex

venn_diagrams.pdf: ../../*.tex ../../trademark.pdf *.tex venn_diagrams.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode venn_diagrams.tex
	../../latexrun venn_diagrams.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: venn_diagrams.pdf
	open venn_diagrams.pdf

o:
	make open

terms: venn_diagrams.tex
	bbk_sheet -mode ts venn_diagrams.tex

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
	rm -f venn_diagrams.*
	make reset
	bbk_sheet -mode mk > Makefile