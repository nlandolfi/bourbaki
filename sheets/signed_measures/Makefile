# generated automatically by bbk_sheet

.PHONY: make o spell s clean remake

make:
	make graph.pdf
	make signed_measures.pdf

sheet.tex: sheet.lit
	lit -in sheet.lit -out sheet.tex

signed_measures.tex: sheet.tex
	bbk_sheet -mode c -in sheet.tex > signed_measures.tex

signed_measures.pdf: ../../*.tex ../../trademark.pdf *.tex signed_measures.tex graph.pdf
	# pdflatex --file-line-error -interaction=nonstopmode signed_measures.tex
	../../latexrun signed_measures.tex
	make terms

graph.csv: sheet.tex
	bbk_sheet -mode g -in sheet.tex > graph.csv

graph.graphviz: graph.csv
	bbk graph --csv graph.csv --tmpl ../../graph/graph.tmpl > graph.graphviz

graph.pdf: graph.graphviz
	dot graph.graphviz -o graph.pdf -T pdf

open: signed_measures.pdf
	open signed_measures.pdf

o:
	make open

terms: signed_measures.tex
	bbk_sheet -mode ts signed_measures.tex

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
	rm -f signed_measures.*
	make reset
	bbk_sheet -mode mk > Makefile